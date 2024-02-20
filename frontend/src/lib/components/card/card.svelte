<script lang="ts">
	import { cn } from '$lib/utils';
	import { type PlayingCard, getCardName, sizes } from '.';

	export let card: PlayingCard;
	export let size: keyof typeof sizes = 'md';
	export let playerBorderColor = '';

	const cardName = getCardName(card);
	const cardPath = `/cards/${cardName}.svg`;
	const width = sizes[size];

	const cardClass = cn([
		{
			[`border-${playerBorderColor} border-4`]: playerBorderColor
		},
		`w-[${width}]`
	]);
</script>

<img style:width class={cardClass} src={cardPath} alt={cardName} />
