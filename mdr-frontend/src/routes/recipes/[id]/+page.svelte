<script>
    import { onMount } from 'svelte';
    import { page } from '$app/stores';

    let recipe = null;
    let error = null;
    let loading = true;
    
    $: id = $page.params.id;

    onMount(async () => {
        try {
            const response = await fetch(`${import.meta.env.VITE_BACKEND_API_URL}/recipes/${id}`);
            if (!response.ok) {
                throw new Error('Failed to fetch recipe');
            }
            recipe = await response.json();
        } catch (e) {
            error = e.message;
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
        <h1>{recipe.title}</h1>
        <p><strong>Created Date:</strong> {recipe.created_date}</p>
        <p>{recipe.content}</p>
    </article>
{/if}

<style>
    article {
        padding: 2rem;
    }
</style>
