<script>
	import Navbar from '$lib/components/navbar.svelte';
</script>

<div class="flex h-dvh max-h-screen flex-col bg-background px-8">
	<Navbar />
	<slot />
</div>
