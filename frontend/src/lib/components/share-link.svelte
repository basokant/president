<script lang="ts">
	import { Button } from '$lib/components/ui/button';
	import { Copy } from 'lucide-svelte';

	export let link: string;
</script>

<div class="flex h-10 flex-1 items-center justify-center">
	<div class="flex h-full flex-1 items-center border-2 border-border px-3 py-1 align-middle">
		{link}
	</div>
	<Button on:click={() => navigator.clipboard.writeText(link)} variant="filled" class="size-10">
		<Copy strokeWidth={2.5} />
	</Button>
</div>
