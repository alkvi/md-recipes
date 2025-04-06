<script lang="ts">
    import { onMount } from 'svelte';
    import { page } from '$app/stores';
    import EditableMarkdown from '$lib/components/EditableMarkdown.svelte';
    import type { Recipe } from '$lib/types/recipe';

    let recipe: Recipe | null = null;
    let error: string | null = null;
    let loading = true;
    let isEditing = false;
    let originalContent = '';
    
    $: id = $page.params.id;

    function stripPreamble(content: string): string {
        const preamble_regex = /^(---|\+\+\+)[\s\S]+?\1/;
        return content.replace(preamble_regex, '').trim();
    }

    function handleSave(event: CustomEvent<{ content: string }>): void {
        if (recipe) {
            const editedContent = event.detail.content;            
            recipe.content = editedContent;
            
            // Save changes to the backend
            fetch(`${import.meta.env.VITE_BACKEND_API_URL}/recipes/${id}`, {
                method: 'PUT',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify(recipe)
            })
            .then(response => {
                if (!response.ok) {
                    throw new Error('Failed to save recipe');
                }
                return response.json();
            })
            .then((updatedRecipe: Recipe) => {
                recipe = updatedRecipe;
                originalContent = recipe.content;
                recipe.content = stripPreamble(recipe.content);
            })
            .catch((e: Error) => {
                error = e.message;
            });
        }
    }

    onMount(async () => {
        try {
            const response = await fetch(`${import.meta.env.VITE_BACKEND_API_URL}/recipes/${id}`);
            if (!response.ok) {
                throw new Error('Failed to fetch recipe');
            }
            recipe = await response.json() as Recipe;
            if (recipe) {
                originalContent = recipe.content;
                recipe.content = stripPreamble(recipe.content);
            }
        } catch (e) {
            error = e instanceof Error ? e.message : 'An unknown error occurred';
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
        <h1>{recipe?.title}</h1>
        <p><strong>Created Date:</strong> {recipe?.created_date}</p>
        <hr>
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
</style>
