import type { Profile } from '@/lib/core/profile/Profile';
import { State } from '@/lib/state/State';

export interface IProfilesState {
	profiles: Map<string, Profile>;
}

export class ProfileState extends State<IProfilesState> {
	public profiles = new Map();

	public get data() {
		return {
			profiles: this.profiles,
		};
	}

	public appendProfile(prof: Profile) {
		this.profiles.set(prof.id, prof);
		return this;
	}
}
