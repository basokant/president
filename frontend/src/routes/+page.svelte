<script>
	import { Button } from '$lib/components/ui/button';
	import { Club, Diamond, Heart, HelpCircleIcon, PlusCircle, Spade } from 'lucide-svelte';
</script>

<div class="flex h-full w-full flex-col items-center justify-center gap-7">
	<div class="flex flex-col items-center justify-center gap-0.5">
		<div class="flex w-72 justify-between">
			<Heart strokeWidth={3} size={30} class="text-primary" />
			<Club strokeWidth={3} size={30} />
			<Diamond strokeWidth={3} size={30} class="text-primary" />
			<Spade strokeWidth={3} size={30} />
		</div>
		<h1 class="text-6xl font-bold">President</h1>
		<h2 class="text-2xl font-medium">The Ultimate Social Card Game</h2>
	</div>
	<div class="space-x-2">
		<Button variant="outlinePrimary">
			<HelpCircleIcon />
			<span>How to Play</span>
		</Button>
		<Button>
			<PlusCircle />
			<span>Create Game</span>
		</Button>
	</div>
</div>
