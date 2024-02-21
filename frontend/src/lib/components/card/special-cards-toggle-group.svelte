<script lang="ts">
	import { ToggleGroup } from 'bits-ui';
	import { type PlayingCard, getCardName, type CardName } from '.';
	import Card from './card.svelte';
	import { timeline } from 'motion';
	import { draw, fade } from 'svelte/transition';

	export let selectedCards: CardName[] = [];

	const specialCards: PlayingCard[] = [
		'joker',
		{
			rank: 'jack',
			suit: 'hearts'
		},
		{
			rank: 'queen',
			suit: 'spades',
			name: 'B**** Queen'
		},
		{
			rank: 'king',
			suit: 'hearts',
			name: 'Suicide King'
		}
	];
</script>

<ToggleGroup.Root class="flex flex-row gap-10" type="multiple" bind:value={selectedCards}>
	{#each specialCards as card, i (i)}
		{@const cardName = getCardName(card)}
		<ToggleGroup.Item
			value={cardName}
			class="group relative inline-flex focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2"
		>
			{#if selectedCards.includes(cardName)}
				<svg
					xmlns="http://www.w3.org/2000/svg"
					width="24"
					height="24"
					viewBox="0 0 24 24"
					fill="none"
					stroke="currentColor"
					stroke-width="2"
					stroke-linecap="round"
					stroke-linejoin="round"
					class="lucide lucide-users-round group absolute -left-3 -top-3 z-10 size-6 rounded-full border-4 border-foreground bg-foreground text-background"
				>
					<path out:fade={{ duration: 100 }} in:draw={{ duration: 200 }} d="M 20 6 9 17 l -5 -5" />
				</svg>
			{/if}
			<Card {card} />
		</ToggleGroup.Item>
	{/each}
</ToggleGroup.Root>
