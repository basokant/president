import { type Player, colours, type Colour } from '$lib/types/player';

export function getAvailableColours(
	players: Player[],
	initialColour: Colour | undefined
): Colour[] {
	const usedColours = new Set<string>();
	players.forEach((p) => usedColours.add(p.colour));
	return colours.filter((c) => !usedColours.has(c) || c === initialColour);
}
