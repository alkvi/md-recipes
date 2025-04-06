<script lang="ts">
    import RecipeCard from './RecipeCard.svelte';
    // placeholder for recipes; these should be acquired from db
    const recipe_list: string[] = ['recipe1', 'recipe2', 'recipe3',
    'recipe4', 'recipe5', 'recipe6',
    'recipe7', 'recipe8', 'recipe9',
    'recipe10', 'recipe11', 'recipe12',
    'recipe13', 'recipe14', 'recipe15'];

    import { createScrollArea, melt } from '@melt-ui/svelte';
    const {
        elements: {
        root,
        content,
        viewport,
        corner,
        scrollbarY,
        thumbY,
        thumbX,
        scrollbarX,
        },
    } = createScrollArea({
            type: 'hover',
            dir: 'ltr',
    });

</script>

<div use:melt={$root} class="grid-root">
    <div use:melt={$viewport} class="grid-viewport">
        <div use:melt={$content}>

            <div class="grid-content">

                {#each recipe_list as recipe}
                <RecipeCard>
                    <p slot="card_title">{recipe}</p>
                    <p>test content</p>
                </RecipeCard>
                {/each}
            </div>

        </div>
    </div>
    <div use:melt={$scrollbarY} class="grid-scrollbar-y">
        <div use:melt={$thumbY} class="grid-thumb-y"/>
    <div/>
    </div>
    <div use:melt={$scrollbarX} class="grid-scrollbar-x">
        <div use:melt={$thumbX} class="grid-thumb-x"/>
    </div>
    <div use:melt={$corner} />
</div>

<style>

    .grid-root {
        overflow: hidden; 
        position: relative; 
        border-radius: 0.375rem; 
        border-width: 1px; 
        width: 58rem; 
        height: 30rem; 
        background-color: #ffffff; 
        box-shadow: 0 10px 15px -3px rgba(0, 0, 0, 0.1), 0 4px 6px -2px rgba(0, 0, 0, 0.05);
    }

    .grid-viewport {
        width: 100%;
        height: 100%;
        border-radius: 0.375rem; 
    }

    .grid-content {
		display: grid;
        grid-template-columns: repeat(3, 1fr);
        grid-auto-rows: auto;
        grid-gap: 1rem;
        padding: 2rem;
	}

    .grid-scrollbar-y {
        border: 2px solid #3e3939;
        background-color: rgb(251 113 133);
        display: flex; 
        padding: 1px; 
        border-left-width: 1px; 
        width: 0.625rem; 
        height: 100%; 
        transition-property: color, background-color, border-color, text-decoration-color, fill, stroke;
        transition-timing-function: cubic-bezier(0.4, 0, 0.2, 1);
        transition-duration: 300ms; 
        touch-action: none; 
        user-select: none; 
    }

    .grid-scrollbar-x {
        display: flex; 
        padding: 1px; 
        border-top-width: 1px; 
        width: 100%; 
        height: 0.625rem; 
        touch-action: none; 
        user-select: none; 
    }

    .grid-thumb-y {
        background-color: rgb(254 205 211);
        position: relative; 
        flex: 1 1 0%; 
        border-radius: 9999px; 
    }

    .grid-thumb.x {
        position: relative; 
        border-radius: 9999px; 
    }

</style>
