<script>
	import '../app.pcss';
</script>

<main class="mb-2 h-full w-full rounded-2xl font-medium text-slate-100">
	<slot />
</main>

<style>
	main {
		font-family: Roboto, sans-serif;
		background-color: rgba(48, 52, 70, 0.87);
		border: solid 2px hsl(233deg 25% 40% / 20%);
		-ms-overflow-style: none; /* IE and Edge */
		scrollbar-width: none; /* Firefox */
	}
	main::-webkit-scrollbar {
		display: none;
	}
</style>
