<script lang="ts">
    import { onMount } from 'svelte';
    import { page } from '$app/stores';
    import EditableMarkdown from '$lib/components/EditableMarkdown.svelte';
    import { createSeparator, melt } from '@melt-ui/svelte';
    import type { Recipe } from '$lib/types/recipe';
    import { stripExtension } from '$lib/utils/string';

    let recipe: Recipe | null = null;
    let error: string | null = null;
    let loading = true;
    let isEditing = false;
    let originalContent = '';
    let originalTitle = '';
    
    $: id = $page.params.id;

    // Decorative separator via melt
    const {
        elements: { root: separator },
    } = createSeparator({
        orientation: 'horizontal',
        decorative: true,
    });

    // Handles editing toggle via the title
    // which is a button styled as h1
    function toggleEditing(event: KeyboardEvent | MouseEvent) {
        if (event.type === 'keydown' && (event as KeyboardEvent).key !== 'Enter' && (event as KeyboardEvent).key !== ' ') {
            return;
        }
        isEditing = true;
    }

    // Strip the metadata preamble in markdown content
    function stripPreamble(content: string): string {
        const preamble_regex = /^(---|\+\+\+)[\s\S]+?\1/;
        return content.replace(preamble_regex, '').trim();
    }

    // Save recipe
    function handleSave(detail: { content: string; renderedContent: string }) {
        if (recipe) {
            console.log('Saving recipe...');
            const editedContent = detail.content;
            recipe.content = editedContent;
            
            // Save changes to the backend
            fetch(`${import.meta.env.VITE_BACKEND_API_URL}/recipes/${id}`, {
                method: 'PUT',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({
                    ...recipe,
                    title: recipe.title // Include the potentially updated title
                })
            })
            .then(response => {
                if (!response.ok) {
                    throw new Error('Failed to save recipe');
                }
                return response.json();
            })
            .then(updatedRecipe => {
                recipe = updatedRecipe;
                if (recipe) {
                    originalContent = recipe.content;
                    recipe.content = stripPreamble(recipe.content);
                }
            })
            .catch(e => {
                error = e instanceof Error ? e.message : 'Unknown error';
            });
        }
    }

    // User pressed cancel editing
    function handleCancel() {
        if (recipe) {
            console.log('Cancelling edit recipe...');
            // Revert both content and title to their original values
            recipe.content = stripPreamble(originalContent);
            recipe.title = stripExtension(originalTitle);
        }
        isEditing = false;
    }

    // Fetch recipe when first mounted
    onMount(async () => {
        try {
            const response = await fetch(`${import.meta.env.VITE_BACKEND_API_URL}/recipes/${id}`);
            if (!response.ok) {
                throw new Error('Failed to fetch recipe');
            }
            recipe = await response.json();
            if (recipe) {
                originalContent = recipe.content;
                originalTitle = recipe.title;
                recipe.title = stripExtension(recipe.title);
                recipe.content = stripPreamble(recipe.content);
            }
        } catch (e) {
            error = e instanceof Error ? e.message : 'Unknown error';
        } finally {
            loading = false;
        }
    });
</script>

{#if loading}
    <p>Loading recipe...</p>
{:else if error}
    <p>Error: {error}</p>
{:else}
    <article>
        {#if isEditing && recipe}
            <input 
                type="text"
                bind:value={recipe.title}
                class="title-input"
                placeholder="Recipe title"
            />
        {:else}
            <button 
                class="title-button"
                on:click={toggleEditing}
                on:keydown={toggleEditing}
            >
                <h1>{recipe?.title}</h1>
            </button>
        {/if}
        <p><strong>Last Modified:</strong> {recipe?.modified_date}</p>
        <div use:melt={$separator} class="separator"></div>
        <EditableMarkdown 
            content={isEditing ? originalContent : recipe?.content || ''}
            bind:isEditing
            onSave={handleSave}
	        onCancel={handleCancel}
        />
    </article>
{/if}

<style>
    article {
        padding: 2rem;
    }

    .separator {
        background-color: #ccc;
        margin: 1rem 0;
        height: 1px;
        width: 100%;
    }

    .title-input {
        font-size: 2rem;
        font-weight: bold;
        width: 100%;
        padding: 0.5rem;
        border: 1px solid #ccc;
        border-radius: 4px;
        margin-bottom: 1rem;
    }

    .title-button {
        background: none;
        border: none;
        padding: 0;
        margin: 0;
        text-align: left;
        width: 100%;
        cursor: pointer;
    }

    .title-button:hover h1 {
        background-color: #f5f5f5;
    }

    .title-button:focus-visible h1 {
        outline: 2px solid currentColor;
        outline-offset: 2px;
    }
</style>
