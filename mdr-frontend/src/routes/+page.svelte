<script lang="ts">
    import RecipeGrid from '$lib/components/RecipeGrid.svelte';
</script>
   
<svelte:head>
  <title>MDR - Main view</title>
</svelte:head>

<section>
    <h1>Welcome to MD-Recipes</h1>
    <p>This will be the MAIN SCREEN</p>
</section>

<RecipeGrid/>


<style>

    section {
        display: flex;
        flex-direction: column;
        justify-content: center;
        align-items: center;
        flex: 0.6;
    }

</style>