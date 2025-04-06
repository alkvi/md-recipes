<script lang="ts">
	import { fade } from 'svelte/transition';
	
	const fadeIn = {
		delay: 100,
		duration: 200
	} as const;
	
	const fadeOut = {
		delay: 0,
		duration: 100
	} as const;
	
	let isExpanded = false;
</script>

<nav class:expanded={isExpanded}>
	<button on:click={() => isExpanded = !isExpanded}>
		{#if isExpanded}
			Close
		{:else}
			Open
		{/if}
	</button>
	<ul>
		<li><a href="/">🏠</a> {#if isExpanded}<span in:fade="{fadeIn}" out:fade="{fadeOut}"><a href="/">Home</a></span>{/if}</li>
		<li>📋 {#if isExpanded}<span in:fade="{fadeIn}" out:fade="{fadeOut}">Categories</span>{/if}</li>
		<li><a href="/settings">⚙️</a> {#if isExpanded}<span in:fade="{fadeIn}" out:fade="{fadeOut}"><a href="/settings">Settings</a></span>{/if}</li>
	</ul>
	{#if isExpanded}
	<section in:fade="{fadeIn}" out:fade="{fadeOut}" >
		<hr />
		<ul>
			<h3>
				Tags
			</h3>
			<li>🧾 Italian</li>
			<li>🤖 Dinner</li>
		</ul>
	</section>
	{/if}
</nav>

<style>
	nav {
		height: 100vw;
		background-color: #324754;
		color: #A2B7C4;
		transition: ease-out 200ms;
		width: 60px;
		overflow: hidden;
	}

	.expanded {
		transition: ease-out 200ms;
		width: 200px;
	}
	
	ul {
		list-style: none;
		padding: 20px;
		margin: 0;
	}
	
	li {
		width: 100px;
	}
	
	h3 {
		text-transform: uppercase;
		margin: 0;
		padding: 10px 0;
	}
	
	hr {
	  color: white;
		width: 80%;
	}
	
	button {
		border: none;
		background: none;
		color: #A2B7C4;
		text-transform: uppercase;
	}
</style>