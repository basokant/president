<script lang="ts">
	import { ToggleGroup } from 'bits-ui';
	import { getCardId, type PlayingCard } from '.';
	import Card from './card.svelte';
	import { draw, fade } from 'svelte/transition';

	export let selectedCards: (typeof specialCards)[number][] = [];

	let selectedCardIds: string[] = [];

	$: {
		selectedCards = selectedCardIds.map((i) => specialCards[Number(i)]);
	}

	const specialCards: PlayingCard[] = [
		'joker',
		{
			rank: 'jack',
			suit: 'hearts',
			name: 'One Eyed Jack (x2)'
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
	] as const;
</script>

<ToggleGroup.Root
	class="flex flex-wrap justify-between gap-2"
	type="multiple"
	bind:value={selectedCardIds}
>
	{#each specialCards as card, i (i)}
		{@const cardName = getCardId(card)}
		{@const label = card === 'joker' ? 'Joker (x2)' : card?.name ?? ''}
		<ToggleGroup.Item
			id={cardName}
			value={i.toString()}
			class="group "
			on:click={() => console.log('clicked', cardName)}
			asChild
			let:builder
		>
			<div
				use:builder.action
				{...builder}
				class="relative pb-2 ring-offset-background focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2"
			>
				<Card {card} {label} />
				{#if selectedCardIds.includes(i.toString())}
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
						class="lucide lucide-users-round group absolute -left-3 -top-3 z-10 size-6 rounded-full border-4 border-foreground bg-foreground text-background dark:bg-white"
					>
						<path
							out:fade={{ duration: 100 }}
							in:draw={{ duration: 200 }}
							d="M 20 6 9 17 l -5 -5"
						/>
					</svg>
				{/if}
			</div>
		</ToggleGroup.Item>
	{/each}
</ToggleGroup.Root>
