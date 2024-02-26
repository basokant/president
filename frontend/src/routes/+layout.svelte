<script>
	import '../app.pcss';
	import '@fontsource-variable/dm-sans';
	import { ModeWatcher } from 'mode-watcher';
	import { Toaster } from 'svelte-sonner';
</script>

<main class="relative h-dvh w-screen">
	<ModeWatcher />
	<Toaster />
	<slot />
</main>
