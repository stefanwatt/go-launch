<script>
	import { LaunchApp } from '$lib/wailsjs/go/main/App';
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
		LaunchApp(desktopEntry.Exec);
	}}
	style="background-color: {selected ? '#aad1fa' : 'rgb(203, 213, 225)'}"
	class="flex h-32 w-60 cursor-pointer items-center rounded-md px-2 text-slate-800"
>
	<div class="w-full">
		<div class="flex justify-center">
			<img class="h-20 w-20" src={`app-icons/${desktopEntry.Icon}`} alt="icon" />
		</div>
		<div class="flex justify-center">
			<div class="flex items-center px-2">
				{desktopEntry.Name}
			</div>
		</div>
	</div>
</div>
<!-- content here -->
