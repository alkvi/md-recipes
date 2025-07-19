<script lang="ts">
    import type { Snippet } from 'svelte';

    export const content: string = "test";
    let { children, card_title }: { children?: Snippet; card_title?: Snippet } = $props();
</script>

<div class="main">
    <h1>
        {@render card_title?.()}
    </h1>
    {@render children?.()}
</div>

<style>

    .main {
        border: 2px solid #3e3939;
        border-radius: 4px;
        padding: .5rem;
    }

</style>
