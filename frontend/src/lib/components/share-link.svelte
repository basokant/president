<script lang="ts">
	import { Button } from '$lib/components/ui/button';
	import { Copy } from 'lucide-svelte';
	import { toast } from 'svelte-sonner';

	export let link: string;

	function copyLink() {
		navigator.clipboard.writeText(link);
		toast.success('Game link copied to clipboard', {
			description: 'Send it to your friends to invite them.'
		});
	}
</script>

<div class="flex min-h-10 flex-1 items-center justify-center">
	<div
		class="flex h-full flex-1 items-center overflow-hidden text-ellipsis border-2 border-border px-3 py-1 align-middle"
	>
		{link}
	</div>
	<Button on:click={copyLink} variant="filled" class="size-9 p-2 lg:size-10 lg:p-3">
		<Copy strokeWidth={2.5} />
	</Button>
</div>
