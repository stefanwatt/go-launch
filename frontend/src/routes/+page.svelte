<script>
	import { GetDesktopEntries } from '$lib/wailsjs/go/main/App';
	import { onMount } from 'svelte';
	import DesktopEntryComponent from './DesktopEntry.svelte';
	import { setup } from './resize-observer';
	import { onKeyPress, onInputKeypress } from './keyboard';
	import {
		promptInput,
		keyboardNavigationInput,
		selectionPosition,
		selectedEntry,
		desktopEntries,
		searchTerm,
		searchResults
	} from './store';
	import SearchIcon from './SearchIcon.svelte';

	onMount(async () => {
		setup();
		$desktopEntries = await GetDesktopEntries();
	});

	/** @param {HTMLInputElement} node */
	function setupKeyboadrNavigation(node) {
		onKeyPress(node);
	}

	selectionPosition.subscribe((pos) => {
		if ($promptInput && !pos && document.activeElement !== $promptInput) $promptInput.focus();
		else if ($keyboardNavigationInput && document.activeElement !== $keyboardNavigationInput)
			$keyboardNavigationInput.focus();
	});
</script>

<input
	bind:this={$keyboardNavigationInput}
	use:setupKeyboadrNavigation
	type="text"
	class="absolute h-0 w-0 overflow-hidden"
/>
<div id="main" class="h-full">
	<div
		id="input-container"
		class="flex flex-row justify-center rounded-b-lg border-b border-fuchsia-200 bg-transparent px-12"
	>
		<!-- svelte-ignore a11y-autofocus -->
		<div class="flex w-full">
			<div class="mr-4 flex h-full w-8 items-center text-slate-100">
				<SearchIcon></SearchIcon>
			</div>
			<input
				bind:this={$promptInput}
				autofocus
				on:keyup={onInputKeypress}
				bind:value={$searchTerm}
				class="h-20 w-full rounded-full border-none bg-transparent text-2xl text-slate-100 outline-none active:border-none"
				placeholder="Search...."
				type="text"
			/>
		</div>
	</div>
	{#if $searchResults.some((row) => row.length !== 0)}
		<div class="mt-2 flex h-full w-full justify-center py-2">
			<div class="grid h-full w-full grid-cols-4 rounded-3xl bg-transparent p-2">
				{#each $searchResults as _, row}
					{#each $searchResults[row] as desktopEntry, col}
						<div
							on:mouseenter={() => {
								console.log('mouseenter');
								$selectionPosition = { row, col };
							}}
							on:mouseleave={() => {
								console.log('mouseleave');
								$selectionPosition = null;
							}}
							class={`m-4 h-32 w-56 col-start-${col} row-start-${row} col-span-1 row-span-1 m-1`}
						>
							<DesktopEntryComponent
								selected={desktopEntry.Exec === $selectedEntry?.Exec}
								{desktopEntry}
							></DesktopEntryComponent>
						</div>
					{/each}
				{/each}
			</div>
		</div>
	{/if}
</div>

<style>
	#input-container {
		border-bottom: solid 2px hsl(233deg 25% 40% / 20%);
	}
</style>
