<script lang="ts">
    import markdownit from 'markdown-it';
    import DOMPurify from 'dompurify';
    import { createSeparator, melt } from '@melt-ui/svelte';

    // Props
	let { onSave, title = '', content = '', modifiedDate = '', isEditing = false } = $props<{
        onSave: (detail: { content: string; title: string }) => void;
        title?: string;
        content?: string;
        modifiedDate?: string;
        isEditing?: boolean;
    }>();
    let draftTitle = $state('');
    let draftContent = $state('');

    // Decorative separator via melt
    const {
        elements: { root: separator },
    } = createSeparator({
        orientation: 'horizontal',
        decorative: true,
    });

    // Initialize markdown-it renderer
    const md = markdownit();

    // Override image renderer
    const defaultRender = md.renderer.rules.image || function (tokens, idx, options, env, self) {
        return self.renderToken(tokens, idx, options);
    };

    // Custom rule for stripping preamble
    md.core.ruler.before('normalize', 'strip-preamble', (state) => {
        const preambleRegex = /^(---|\+\+\+)[\s\S]+?\1/;
        state.src = state.src.replace(preambleRegex, '').trim();
        return true;
    });

    // Custom rule for replacing local image path with backend path
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

    // Initialize drafts on edit toggle
    $effect(() => {
        if (isEditing) {
            draftTitle = title;
            draftContent = content;
        }
    });

    function handleSave(): void {
        console.log("Saving recipe")
        title = draftTitle;
        content = draftContent;
        isEditing = false;
        onSave?.({ title: draftTitle, content: draftContent });
    }

    function handleCancel(): void {
        console.log("Cancel edit recipe")
        isEditing = false;
    }
</script>

{#if isEditing}        
    <input
        type="text"
        bind:value={draftTitle}
        class="title-input"
        placeholder="Recipe title"
        onkeydown={(e) => {
            if (e.key === 'Enter') handleSave();
            if (e.key === 'Escape') handleCancel();
        }}
    />

    <div use:melt={$separator} class="separator"></div>

    <textarea 
        bind:value={draftContent}
        rows="20"
        placeholder=""
    ></textarea>
    <div class="edit-actions">
        <button onclick={handleSave}>Save</button>
        <button onclick={handleCancel}>Cancel</button>
    </div>

    <p><strong>Last Modified:</strong> {modifiedDate}</p>
{:else}
    <button
        class="title-button"
        onclick={() => isEditing = true} 
        onkeydown={(e) => e.key === 'Enter' && (isEditing = true)}
        aria-label="Edit title"
    >
        <h1>{title}</h1>
    </button>

    <div use:melt={$separator} class="separator"></div>

    <div 
        class="content" 
        onclick={() => isEditing = true} 
        onkeydown={(e) => e.key === 'Enter' && (isEditing = true)}
        role="button" 
        tabindex="0"
    >
        {@html DOMPurify.sanitize(md.render(content))}
    </div>

    <p><strong>Last Modified:</strong> {modifiedDate}</p>
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

    .title-button {
        font-size: 1.5rem;
        background: none;
        border: none;
        margin: 0;
        text-align: left;
        width: 100%;
        color: black;
        cursor: pointer;
        padding-left: 1rem;
        border-radius: 4px;
    }

    .title-button h1 {
        margin: 0;
        font-size: 2rem;
        font-weight: bold;
    }

    .title-button:hover {
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

    .title-input {
        font-size: 2rem;
        font-weight: bold;
        width: 100%;
        padding: 0.5rem;
        border: 1px solid #ccc;
        border-radius: 4px;
        margin-bottom: 0rem;
    }

    .separator {
        background-color: #ccc;
        margin: 1rem 0;
        height: 1px;
        width: 100%;
    }
</style> 