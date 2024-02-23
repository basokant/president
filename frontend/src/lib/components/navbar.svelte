<script>
	import { HelpCircle, Club, Diamond, Heart, Spade } from 'lucide-svelte';
	import { Button } from '$lib/components/ui/button';
	import Separator from './ui/separator/separator.svelte';
</script>

<nav class="flex flex-row items-center justify-between py-2 pt-3">
	<a href="/" class="text-2xl font-bold">President</a>
	<div class="flex w-72 justify-between">
		<Heart strokeWidth={3} size={20} class="text-primary" />
		<Club strokeWidth={3} size={20} />
		<Diamond strokeWidth={3} size={20} class=" text-primary " />
		<Spade strokeWidth={3} size={20} />
	</div>
	<Button variant="outlinePrimary">
		<HelpCircle />
		<span>How to Play</span>
	</Button>
</nav>
<Separator orientation="horizontal" />
