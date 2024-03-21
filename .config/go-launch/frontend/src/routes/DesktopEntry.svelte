<script>
	import { LaunchApp } from '$lib/go-launch/app.iframe';
	/** @type {App.DesktopEntry} */
	export let desktopEntry;
	/** @type {boolean} */
	export let selected;

	/**@param{HTMLDivElement} node*/
	function autoAdjustFontSize(node) {
		import('fitty').then((module) => {
			const fitty = module.default;
			fitty(node, { minSize: 10, maxSize: 16, multiLine: true });
		});
	}
</script>

<!-- svelte-ignore a11y-click-events-have-key-events -->
<!-- eslint-disable-next-line svelte/valid-compile -->
<!-- svelte-ignore a11y-no-static-element-interactions -->
<div
	on:click={() => {
		LaunchApp(desktopEntry.Id);
	}}
	style={selected ? 'background-color: #838ba7; color: #1e293b' : 'background-color: transparent'}
	class="flex h-full w-full cursor-pointer items-center rounded-2xl px-2 text-slate-100"
>
	<div class="w-full">
		<div class="flex justify-center">
			<img class="h-16 w-16" src={`app-icons/${desktopEntry.Icon}`} alt="icon" />
		</div>
		<div class="mt-4 flex h-4 justify-center">
			<div use:autoAdjustFontSize class="flex items-center justify-center px-2">
				{desktopEntry.Name}
			</div>
		</div>
	</div>
</div>
