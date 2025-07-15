<script lang="ts">
    import RecipeGrid from '$lib/components/RecipeGrid.svelte';
    import { backendConfig } from '$lib/stores/backendConfig';
</script>
   
<svelte:head>
  <title>MDR - Main view</title>
</svelte:head>

<section>
    <h1>Welcome to MD-Recipes</h1>
    <p>This will be the MAIN SCREEN</p>
    {#if $backendConfig}
	    <p>Folder path: {$backendConfig.folderPath}</p>
    {/if}
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