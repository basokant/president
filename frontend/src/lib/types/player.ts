import { z } from 'zod';

export type Player = {
	id: number;
	name: string;
	isUser?: boolean;
	isHost?: boolean;
	isComputer?: boolean;
	colour: Colour;
};

export const colours = [
	'red',
	'orange',
	'yellow',
	'lime',
	'emerald',
	'cyan',
	'blue',
	'indigo',
	'purple',
	'pink'
] as const;

export type Colour = (typeof colours)[number];

export function createPlayerSchema(availableColours: Colour[]) {
	const playerSchema = z.object({
		name: z.string().min(3).max(15),
		colour: z.enum(colours).refine((c) => availableColours.includes(c))
	});

	return playerSchema;
}

export const gameSchema = z.object({
	name: z.string().min(3).max(15),
	colour: z.enum(colours)
});
