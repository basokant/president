<script lang="ts">
	import { Button } from '$lib/components/ui/button';
	import { Club, Diamond, Heart, HelpCircleIcon, PlusCircle, Spade } from 'lucide-svelte';
</script>

<svelte:head>
	<title>President | The Ultimate Social Card Game</title>
</svelte:head>

<img class="absolute -z-10" src="/pattern.svg" alt="playing card symbol pattern" />
<img
	class="absolute bottom-0 right-0 -z-10 rotate-180"
	src="/pattern.svg"
	alt="playing card symbol pattern"
/>
<div class="z-10 flex h-full w-full flex-col items-center justify-center gap-7">
	<div class="flex flex-col items-center justify-center gap-0.5">
		<div class="flex w-72 justify-between">
			<Heart strokeWidth={3} size={30} class="animate-bounce text-primary" />
			<Club strokeWidth={3} size={30} class="animate-[bounce_1.2s_infinite]" />
			<Diamond strokeWidth={3} size={30} class="animate-[bounce_1.4s_infinite] text-primary " />
			<Spade strokeWidth={3} size={30} class="animate-[bounce_1.6s_infinite]" />
		</div>
		<h1 class="text-6xl font-bold">President</h1>
		<h2 class="text-2xl font-medium">The Ultimate Social Card Game</h2>
	</div>
	<div class="space-x-4">
		<Button variant="outlinePrimary">
			<HelpCircleIcon />
			<span>How to Play</span>
		</Button>
		<Button href="/game/15">
			<PlusCircle />
			<span>Create Game</span>
		</Button>
	</div>
</div>
