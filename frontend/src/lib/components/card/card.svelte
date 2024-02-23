<script lang="ts">
	import { cn } from '$lib/utils';
	import { type PlayingCard, getCardId, sizes } from '.';
	import * as Tooltip from '$lib/components/ui/tooltip';
	import { Label } from '../ui/label';

	export let card: PlayingCard;
	export let size: keyof typeof sizes = 'md';
	export let playerBorderColor = '';
	export let label: string;

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
	<Tooltip.Trigger class="flex flex-col items-center gap-2">
		<img style:width class={cardClass} src={cardPath} alt={cardName} />
		<Label>
			{label}
		</Label>
	</Tooltip.Trigger>
	<Tooltip.Content>
		{cardName}
	</Tooltip.Content>
</Tooltip.Root>
