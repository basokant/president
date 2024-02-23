<script lang="ts">
	import { cn } from '$lib/utils';
	import { type PlayingCard, getCardId, sizes } from '.';
	import * as Tooltip from '$lib/components/ui/tooltip';

	export let card: PlayingCard;
	export let size: keyof typeof sizes = 'md';
	export let playerBorderColor = '';

	$: cardName = getCardId(card);
	$: cardPath = `/cards/${cardName}.svg`;
	$: width = sizes[size];

	$: cardClass = cn([
		{
			[`border-${playerBorderColor} border-4`]: playerBorderColor
		},
		`w-[${width}]`
	]);
</script>

<Tooltip.Root openDelay={100}>
	<Tooltip.Trigger>
		<img style:width class={cardClass} src={cardPath} alt={cardName} />
	</Tooltip.Trigger>
	<Tooltip.Content>
		{cardName}
	</Tooltip.Content>
</Tooltip.Root>
