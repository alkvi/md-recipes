<script lang="ts">
    import type { Snippet } from 'svelte';

    let { children, card_title, card_image }: { 
        children?: Snippet; 
        card_title?: Snippet;
        card_image?: Snippet;
    } = $props();
</script>

<div class="main">
    <h1>
        {@render card_title?.()}
    </h1>
    {@render card_image?.()}
    {@render children?.()}
</div>

<style>

    .main {
        border: 2px solid #3e3939;
        border-radius: 4px;
        padding: .5rem;
    }

</style>
