<script lang="ts">
    import { createEventDispatcher } from 'svelte';
    import markdownit from 'markdown-it';
    import DOMPurify from 'dompurify';

    const dispatch = createEventDispatcher<{
        save: { content: string; renderedContent: string };
        cancel: void;
    }>();
    const md = markdownit();

    export let content = '';
    export let isEditing = false;

    let rawContent = '';

    $: if (isEditing) {
        rawContent = content;
    }

    function handleSave(): void {
        content = rawContent;
        const renderedContent = DOMPurify.sanitize(md.render(content));
        dispatch('save', { content, renderedContent });
        isEditing = false;
    }

    function handleCancel(): void {
        isEditing = false;
        dispatch('cancel');
    }
</script>

{#if isEditing}
    <div class="edit-container">
        <textarea 
            bind:value={rawContent}
            rows="20"
            placeholder=""
        ></textarea>
        <div class="edit-actions">
            <button on:click={handleSave}>Save</button>
            <button on:click={handleCancel}>Cancel</button>
        </div>
    </div>
{:else}
    <div class="content" on:click={() => isEditing = true}>
        {@html DOMPurify.sanitize(md.render(content))}
    </div>
{/if}

<style>
    .content {
        cursor: pointer;
        padding: 1rem;
        border-radius: 4px;
    }
    
    .content:hover {
        background-color: #f5f5f5;
    }
    
    .edit-container {
        width: 100%;
    }
    
    textarea {
        width: 100%;
        padding: 1rem;
        border: 1px solid #ccc;
        border-radius: 4px;
        font-family: inherit;
        font-size: inherit;
        line-height: 1.5;
        margin-bottom: 1rem;
    }
    
    .edit-actions {
        display: flex;
        gap: 1rem;
    }
    
    button {
        padding: 0.5rem 1rem;
        border: none;
        border-radius: 4px;
        cursor: pointer;
        background-color: #4a4a4a;
        color: white;
    }
    
    button:hover {
        background-color: #363636;
    }
</style> 