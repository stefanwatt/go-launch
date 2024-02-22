<script>
	import { GetDesktopEntries, LaunchApp, HideLauncher } from '$lib/wailsjs/go/main/App';
	import { onMount } from 'svelte';
	import DesktopEntryComponent from './DesktopEntry.svelte';
	import { WindowSetSize } from '$lib/wailsjs/runtime/runtime';

	/** @type {App.DesktopEntry[]} */
	let desktopEntries = [];
	let searchTerm = '';

	onMount(async () => {
		desktopEntries = await GetDesktopEntries();
		/** @param {KeyboardEvent} e*/
		var ro = new ResizeObserver((entries) => {
			for (let entry of entries) {
				const cr = entry.contentRect;
				WindowSetSize(cr.width, cr.height);
			}
		});
		const mainElement = document.getElementById('main');
		if (!mainElement) return;
		ro.observe(mainElement);
	});

	window.addEventListener('keyup', (e) => {
		if (!selectedEntry) return;
		const index = searchResults.indexOf(selectedEntry);
		switch (e.key) {
			case 'ArrowUp':
				setSelectedEntry(searchResults[index - 4]);
				break;
			case 'ArrowDown':
				setSelectedEntry(searchResults[index + 4]);
				break;
			case 'ArrowLeft':
				setSelectedEntry(searchResults[index - 1]);
				break;
			case 'ArrowRight':
				setSelectedEntry(searchResults[index + 1]);
				break;
			case 'Enter':
				LaunchApp(selectedEntry.Exec);
				searchTerm = '';
				break;
			case 'Escape':
				searchTerm = '';
				HideLauncher();
				break;
			default:
				break;
		}
	});

	/** @type {App.DesktopEntry[]} */
	$: searchResults =
		searchTerm === ''
			? []
			: desktopEntries.filter((entry) =>
					entry.Name.toLowerCase().includes(searchTerm.toLowerCase())
				);

	$: {
		if (searchResults.length === 1) {
			setSelectedEntry(searchResults[0]);
		}
	}
	/** @type {App.DesktopEntry?} */
	let selectedEntry;

	/** @param {KeyboardEvent} e*/
	function startKeyboardNavigation(e) {
		if (e.key === 'Escape') CloseApp();
		if (selectedEntry) return;
		if (e.key !== 'ArrowDown') return;
		setSelectedEntry(searchResults[0]);
	}

	/** @param {App.DesktopEntry} entry*/
	function setSelectedEntry(entry) {
		selectedEntry = entry;
		searchResults = searchResults;
	}
	function resetSelectedEntry() {
		selectedEntry = null;
		searchResults = searchResults;
	}
</script>

<div id="main">
	<div class="flex flex-row justify-center bg-slate-700 px-2">
		<!-- svelte-ignore a11y-autofocus -->
		<input
			autofocus
			on:keydown={startKeyboardNavigation}
			bind:value={searchTerm}
			class="h-20 w-full border-none bg-slate-700 text-xl text-slate-100 outline-none active:border-none"
			placeholder="Search...."
			type="text"
		/>
	</div>
	{#if searchResults.length}
		<div class="flex w-full justify-center bg-slate-600 p-2">
			<div class="flex flex-wrap">
				{#each searchResults.slice(0, 16) as desktopEntry}
					<div
						on:mouseenter={() => {
							setSelectedEntry(desktopEntry);
						}}
						on:mouseleave={() => {
							resetSelectedEntry();
						}}
						class="m-1"
					>
						<DesktopEntryComponent
							selected={desktopEntry.Exec === selectedEntry?.Exec}
							{desktopEntry}
						></DesktopEntryComponent>
					</div>
				{/each}
			</div>
		</div>
	{/if}
</div>
