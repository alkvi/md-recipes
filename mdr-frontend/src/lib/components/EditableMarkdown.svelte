<script lang="ts">
    import markdownit from 'markdown-it';
    import DOMPurify from 'dompurify';

	export let onSave: (detail: { content: string; renderedContent: string }) => void;
	export let onCancel: () => void;
    export let content = '';
    export let isEditing = false;
    let rawContent = '';

    // Initialize markdown-it renderer
    const md = markdownit();

    // Override image renderer
    const defaultRender = md.renderer.rules.image || function (tokens, idx, options, env, self) {
        return self.renderToken(tokens, idx, options);
    };

    // Add a rule to image renderer, so that
    // local image path is replaced with backend path
    // TODO: get endpoint from conf
    md.renderer.rules.image = function (tokens, idx, options, env, self) {
        const token = tokens[idx];
        const srcIndex = token.attrIndex('src');

        if (srcIndex >= 0) {
            const srcAttr = token.attrs![srcIndex];
            const originalSrc = srcAttr[1];

            if (originalSrc.startsWith('../images/')) {
            // Rewrite path to backend endpoint
            srcAttr[1] = originalSrc.replace('../images/', 'http://localhost:3000/images/');
            }
        }

        return defaultRender(tokens, idx, options, env, self);
    };

    $: if (isEditing) {
        rawContent = content;
    }

    function handleSave(): void {
        content = rawContent;
        const renderedContent = DOMPurify.sanitize(md.render(content));
        onSave?.({ content: rawContent, renderedContent });
        isEditing = false;
    }

    function handleCancel(): void {
        isEditing = false;
        onCancel?.();
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
    <div 
        class="content" 
        on:click={() => isEditing = true} 
        on:keydown={(e) => e.key === 'Enter' && (isEditing = true)}
        role="button" 
        tabindex="0"
    >
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