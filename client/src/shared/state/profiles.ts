import type { Profile } from '@/lib/core/profile/Profile';
import { State } from '@/lib/state/State';

export interface ProfileStateData {
	profiles: Map<string, Profile>;
	activeProfile?: Profile;
}

export class ProfileState extends State<ProfileStateData> {
	private profiles = new Map<string, Profile>();
	private activeProfile?: Profile;

	public get data() {
		return {
			profiles: this.profiles,
			activeProfile: this.activeProfile,
		};
	}

	public addProfile(profile: Profile) {
		this.profiles.set(profile.id, profile);
		return this;
	}

	public setActiveProfile(profile: Profile) {
		this.activeProfile = profile;
		return this;
	}
}
