<script>
	import { flip } from 'svelte/animate';
	import { FuzzyFindDesktopEntry } from '$lib/wailsjs/go/main/App';
	import DesktopEntryComponent from './DesktopEntry.svelte';
	import { setupResizeObserver } from './resize-observer';
	import { onKeyPress, onInputKeypress } from './keyboard';
	import {
		promptInput,
		keyboardNavigationInput,
		selectionPosition,
		selectedEntry,
		searchTerm,
		searchResults
	} from './store';
	import SearchIcon from './SearchIcon.svelte';
	import { slide, fade } from 'svelte/transition';
	import { quintOut } from 'svelte/easing';
	import { onMount } from 'svelte';

	/** @param {HTMLInputElement} node */
	function setupKeyboardNavigation(node) {
		onKeyPress(node);
	}

	onMount(async () => {
		const results = await FuzzyFindDesktopEntry('');
		searchResults.set(results);
	});

	selectionPosition.subscribe(
		/** @param {import('./store').Position} pos*/ (pos) => {
			if ($promptInput && !pos && document.activeElement !== $promptInput) $promptInput.focus();
			else if ($keyboardNavigationInput && document.activeElement !== $keyboardNavigationInput)
				$keyboardNavigationInput.focus();
		}
	);

	searchResults.subscribe(
		/** @param {App.DesktopEntry[][]} newSearchResults*/ (newSearchResults) => {
			console.log(newSearchResults);
			const noResults = newSearchResults.every((row) => row.length === 0);
			const promptInputFocused = document?.activeElement === $promptInput;
			if ($promptInput && !promptInputFocused && noResults) {
				return $promptInput.focus();
			} else if (
				$searchResults?.length &&
				$searchResults[0].filter(/** @param {App.DesktopEntry} entry*/ (entry) => !!entry)
					.length === 1
			) {
				selectionPosition.set({ row: 0, col: 0 });
			}
		}
	);

	let lastSearchTerm = '';
	searchTerm.subscribe(
		/** @param {string} value*/ async (value) => {
			if (value === lastSearchTerm) return;
			lastSearchTerm = value;
			const entries = await FuzzyFindDesktopEntry(value);
			/**@type {App.DesktopEntry[]}*/
			searchResults.set(entries);
		}
	);
</script>

<input
	bind:this={$keyboardNavigationInput}
	use:setupKeyboardNavigation
	type="text"
	class="absolute h-0 w-0 overflow-hidden"
/>
<div use:setupResizeObserver class="h-full">
	<div id="input-container" class="flex flex-row justify-center bg-transparent px-12">
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
				class="h-20 w-full border-none bg-transparent text-2xl text-slate-100 outline-none active:border-none"
				placeholder="Search...."
				type="text"
			/>
		</div>
	</div>
	{#if $searchResults.some(/** @param {App.DesktopEntry[]} row */ (row) => row.length !== 0)}
		<div class="mt-2 flex h-full w-full justify-center py-2">
			<div class="h-full w-full rounded-3xl bg-transparent p-2">
				{#each $searchResults as _, row}
					{#if $searchResults[row]?.some(/** @param {App.DesktopEntry} entry */ (entry) => !!entry)}
						<div transition:slide={{ delay: 0, duration: 300, easing: quintOut }} class="flex">
							{#each $searchResults[row].filter(/**@param {App.DesktopEntry}entry*/ (entry) => !!entry) as desktopEntry, col (desktopEntry.Id)}
								<div
									class="w-1/4"
									in:fade={{ delay: 0, duration: 500, easing: quintOut }}
									animate:flip={{ duration: 500, easing: quintOut }}
								>
									<!-- svelte-ignore a11y-no-static-element-interactions -->
									<div
										on:mouseenter={() => {
											$selectionPosition = { row, col };
										}}
										on:mouseleave={() => {
											$selectionPosition = null;
										}}
										class={`m-4 h-32 w-56 col-start-${col} row-start-${row} col-span-1 row-span-1 m-1`}
									>
										<DesktopEntryComponent
											selected={desktopEntry.Id === $selectedEntry?.Id}
											{desktopEntry}
										></DesktopEntryComponent>
									</div>
								</div>
							{/each}
						</div>
					{/if}
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
