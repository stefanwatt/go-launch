<script>
	import { LaunchApp } from '$lib/wailsjs/go/main/App';
	// TODO: maybe make highlighed entry sexier
	/** @type {App.DesktopEntry} */
	export let desktopEntry;
	/** @type {boolean} */
	export let selected;

	/** @type {any} */
	let icon = null;

	// import(`$lib/app-icons/${desktopEntry.Name.replaceAll(' ', '')}-icon.svg`).then((module) => {
	// 	icon = module.default;
	// });
	/** @type {any} */

	// onMount(async () => {
	// 	iconComponent = await import(`$lib/app-icons/${desktopEntry.Icon}.svelte`);
	// });
</script>

<!-- svelte-ignore a11y-click-events-have-key-events -->
<!-- eslint-disable-next-line svelte/valid-compile -->
<!-- svelte-ignore a11y-no-static-element-interactions -->
<div
	on:click={() => {
		LaunchApp(desktopEntry.Id);
	}}
	style={selected ? 'background-color: #aad1fa; color: #1e293b' : 'background-color: transparent'}
	class="flex h-full w-full cursor-pointer items-center rounded-2xl px-2 text-slate-100"
>
	<div class="w-full">
		<div class="flex justify-center">
			<img class="h-16 w-16" src={`app-icons/${desktopEntry.Icon}`} alt="icon" />
		</div>
		<div class="mt-4 flex h-4 justify-center">
			<div class="flex items-center justify-center px-2">
				{desktopEntry.Name.slice(0, 22)}
			</div>
		</div>
	</div>
</div>
<!-- content here -->
