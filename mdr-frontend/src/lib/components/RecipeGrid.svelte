<script>
    import { onMount } from 'svelte';
    import RecipeCard from './RecipeCard.svelte';

    let error = null;
    let loading = true;
    let recipes = [];

    onMount(async () => {
        try {
            console.log("Fetching recipes");
            const response = await fetch(`${import.meta.env.VITE_BACKEND_API_URL}/recipes`);
            if (!response.ok) {
                throw new Error('Failed to fetch recipes');
            }
            recipes = await response.json();
        } catch (e) {
            error = e.message;
        } finally {
            loading = false;
        }
    });

</script>

<div class="grid-content">

{#if loading}
    <p>Loading recipes...</p>
{:else if error}
    <p>Error: {error}</p>
{:else}
    {#each recipes as recipe}
        <RecipeCard>
            <a slot="card_title" href={`/recipes/${recipe.id}`}>{recipe.title}</a>
            <p>{recipe.content}</p>
        </RecipeCard>
    {/each}
{/if}

</div>


<style>

    .grid-content {
		display: grid;
        grid-template-columns: repeat(3, 1fr);
        grid-auto-rows: auto;
        grid-gap: 1rem;
        padding: 2rem;
	}

</style>
