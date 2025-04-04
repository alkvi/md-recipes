<script>
    import { onMount } from 'svelte';
    import { page } from '$app/stores';
    import markdownit from 'markdown-it';
    import DOMPurify from 'dompurify';

    let recipe = null;
    let error = null;
    let loading = true;
    let isEditing = false;
    let rawContent = '';
    
    $: id = $page.params.id;
    const md = markdownit();
    let recipe_content = "";

    function stripPreamble(content) {
        // Match content between --- or +++ markers at the start of the content
        const preamble_regex = /^(---|\+\+\+)[\s\S]+?\1/;
        return content.replace(preamble_regex, '').trim();
    }

    function toggleEdit() {
        isEditing = !isEditing;
        if (isEditing && recipe) {
            rawContent = recipe.content;
        }
    }

    function handleSave() {
        if (recipe) {
            recipe.content = rawContent;
            const content_stripped = stripPreamble(recipe.content);
            recipe_content = DOMPurify.sanitize(md.render(content_stripped));
            isEditing = false;
            // TODO: Add API call to save changes
        }
    }

    onMount(async () => {
        try {
            const response = await fetch(`${import.meta.env.VITE_BACKEND_API_URL}/recipes/${id}`);
            if (!response.ok) {
                throw new Error('Failed to fetch recipe');
            }
            recipe = await response.json();
            const content_stripped = stripPreamble(recipe.content);
            recipe_content = DOMPurify.sanitize(md.render(content_stripped));
        } catch (e) {
            error = e.message;
        } finally {
            loading = false;
        }
    });
</script>

{#if loading}
    <p>Loading recipe...</p>
{:else if error}
    <p>Error: {error}</p>
{:else}
    <article>
        <h1>{recipe.title}</h1>
        <p><strong>Created Date:</strong> {recipe.created_date}</p>
        <hr>
        {#if isEditing}
            <div class="edit-container">
                <textarea 
                    bind:value={rawContent}
                    rows="20"
                    placeholder=""
                ></textarea>
                <div class="edit-actions">
                    <button on:click={handleSave}>Save</button>
                    <button on:click={toggleEdit}>Cancel</button>
                </div>
            </div>
        {:else}
            <div class="content" on:click={toggleEdit}>
                {@html recipe_content}
            </div>
        {/if}
    </article>
{/if}

<style>
    article {
        padding: 2rem;
    }
    
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
