<script>
	import '../app.pcss';
	import '@fontsource-variable/dm-sans';
	import { ModeWatcher } from 'mode-watcher';
</script>

<main class="relative h-dvh w-screen">
	<ModeWatcher />
	<slot />
</main>
