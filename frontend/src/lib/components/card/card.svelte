<script lang="ts">
	import { cn } from '$lib/utils';
	import { type PlayingCard, getCardName, sizes } from '.';
	import * as Tooltip from '$lib/components/ui/tooltip';

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

<Tooltip.Root openDelay={100}>
	<Tooltip.Trigger>
		<img style:width class={cardClass} src={cardPath} alt={cardName} />
	</Tooltip.Trigger>
	<Tooltip.Content side="bottom">
		{cardName}
	</Tooltip.Content>
</Tooltip.Root>
