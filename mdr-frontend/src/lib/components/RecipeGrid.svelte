<script lang="ts">
    import { onMount } from 'svelte';
    import RecipeCard from './RecipeCard.svelte';
    import type { Recipe } from '$lib/types/recipe';
    import { stripExtension } from '$lib/utils/string';

    let error = $state<string | null>(null);
    let loading = $state(true);
    let recipes = $state<Recipe[]>([]);

    onMount(async (): Promise<void> => {
        try {
            console.log("Fetching recipes");
            const response = await fetch(`${import.meta.env.VITE_BACKEND_API_URL}/recipes`);
            if (!response.ok) {
                throw new Error('Failed to fetch recipes');
            }
            recipes = await response.json();
            recipes = recipes.map(recipe => ({
                ...recipe,
                title: stripExtension(recipe.title),
                image_path: recipe.image_path
                    ? recipe.image_path.replace('../images/', 'http://localhost:3000/images/')
                    : null
            }));
        } catch (e) {
            error = e instanceof Error ? e.message : 'Unknown error';
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
            {#snippet card_title()}
					<a href={`/recipes/${recipe.id}`}>{recipe.title}</a>
			{/snippet}
            {#snippet card_image()}
                {#if recipe.image_path}
                    <img src={recipe.image_path} alt="Recipe preview" class="recipe-image" />
                {/if}
            {/snippet}
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

    .recipe-image {
		width: 100%;
		max-height: 200px;
		object-fit: cover;
		margin: 1rem 0;
		border-radius: 4px;
	}

</style>
