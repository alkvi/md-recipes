<script lang="ts">
    import { onMount } from 'svelte';
    import EditableMarkdown from '$lib/components/EditableMarkdown.svelte';
    import type { Recipe } from '$lib/types/recipe';
    import { stripExtension } from '$lib/utils/string';
    import { page } from '$app/stores';

    const { params } = $props();
    let id = $derived($page.params.id);
    let recipe = $state<Recipe | null>(null);
    let error = $state<string | null>(null);
    let loading = $state(true);
    
    // Save recipe
    function handleSave(updated: { title: string, content: string }) {
        if (!recipe) return;

        // Save changes to the backend
        console.log("Sending PUT request for saving recipe")
        fetch(`${import.meta.env.VITE_BACKEND_API_URL}/recipes/${id}`, {
            method: 'PUT',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({
                ...recipe,
                title: updated.title,
                content: updated.content
            })
        })
        .then(response => {
            if (!response.ok) {
                throw new Error('Failed to save recipe');
            }
            console.log("Recipe update OK")
            return response.json();
        })
    }

    // Fetch recipe when first mounted
    onMount(async () => {
        try {
            const response = await fetch(`${import.meta.env.VITE_BACKEND_API_URL}/recipes/${id}`);
            if (!response.ok) throw new Error('Failed to fetch recipe');
            recipe = await response.json();
            if (recipe) recipe.title = stripExtension(recipe.title);
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
        {#if recipe}
            <EditableMarkdown 
                title={recipe.title}
                content={recipe.content}
                modifiedDate={recipe.modified_date}
                onSave={handleSave}
            />
        {/if}
    </article>
{/if}

<style>
    article {
        padding: 2rem;
    }
</style>
