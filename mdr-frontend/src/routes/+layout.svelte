<script lang="ts">
	import { onMount } from 'svelte';
    import Sidebar from "$lib/components/Sidebar.svelte";
	import { backendConfig as configStore, type BackendConfig } from '$lib/stores/backendConfig';

	// `data` comes from +layout.server.js
	let { children, data } = $props();

	onMount(() => {
		// Set the config into the store after hydration
		configStore.set(data.config);
	});

</script>

<div class="app">
	<div class="whole-page">
		<aside>
			<div class="sidebar-div">
				<Sidebar/>
			</div>
		</aside>
		<div class="main-content">
			{@render children?.()}
		</div>
	</div>
</div>


<style>

    .app {
		display: flex;
		flex-direction: column;
		min-height: 100vh;
	}

	.whole-page {
		display: grid;
		grid-template-columns: auto 10fr;
		grid-template-areas: "sidebar content";
		width: 100%;
	}

	.main-content {
		grid-area: content;
		border: 2px solid #3e3939;
	}

	aside {
		grid-area: sidebar;
		overflow-y: scroll; 
		position: sticky; 
		top: 0rem; 
		border: 2px solid #3e3939;
		overflow-y: scroll;
  		max-height: 100vh;
	}

	.sidebar-div {
		border: 2px solid #cc1c1c;
		padding-top: 0rem;
	}

</style>
