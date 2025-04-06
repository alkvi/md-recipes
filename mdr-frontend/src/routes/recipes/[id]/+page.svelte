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

    const {
        elements: { root: separator },
    } = createSeparator({
        orientation: 'horizontal',
        decorative: true,
    });

    function stripPreamble(content: string): string {
        const preamble_regex = /^(---|\+\+\+)[\s\S]+?\1/;
        return content.replace(preamble_regex, '').trim();
    }

    function handleSave(event: CustomEvent<{ content: string; renderedContent: string }>): void {
        if (recipe) {
            const editedContent = event.detail.content;
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
            <h1 on:click={() => isEditing = true}>{recipe?.title}</h1>
        {/if}
        <p><strong>Last Modified:</strong> {recipe?.modified_date}</p>
        <div use:melt={$separator} class="separator" />
        <EditableMarkdown 
            content={isEditing ? originalContent : recipe?.content || ''}
            bind:isEditing
            on:save={handleSave}
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

    h1 {
        cursor: pointer;
    }

    h1:hover {
        background-color: #f5f5f5;
    }
</style>
