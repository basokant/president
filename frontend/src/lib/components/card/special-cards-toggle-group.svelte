<script lang="ts">
	import { ToggleGroup } from 'bits-ui';
	import { type PlayingCard, getCardName, type CardName } from '.';
	import Card from './card.svelte';
	import { Check } from 'lucide-svelte';

	export let value: CardName[] = [];

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

<ToggleGroup.Root class="flex flex-row gap-10" type="multiple" {value}>
	{#each specialCards as card}
		{@const cardName = getCardName(card)}
		<ToggleGroup.Item
			value={cardName}
			class="group relative inline-flex focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2"
		>
			<Check
				class="absolute -left-3 -top-3 z-10 size-6 rounded-full border-4 border-foreground bg-foreground text-background group-data-[state=off]:hidden"
			/>
			<Card {card} />
		</ToggleGroup.Item>
	{/each}
</ToggleGroup.Root>
