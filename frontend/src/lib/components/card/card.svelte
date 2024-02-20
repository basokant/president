<script lang="ts">
	import { type PlayingCard, getCardName } from '.';
	export let card: PlayingCard;

	const cardName = getCardName(card);
	const cardPath = `/cards/${cardName}.svg`;
</script>

<img src={cardPath} alt={cardName} />
