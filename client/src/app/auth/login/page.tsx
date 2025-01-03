"use client";

import { useEffect, useState } from "react";
import {
	browserLocalPersistence,
	onAuthStateChanged,
	setPersistence,
	signInWithPopup,
	type User,
} from "firebase/auth";
import { auth, provider } from "~/firebase/initializeApp";
import { useRouter } from "next/navigation";

export default function LoginPage() {
	const [user, setUser] = useState<User | null>(null);
	const router = useRouter();

	// 認証状態を監視
	useEffect(() => {
		const unsubscribe = onAuthStateChanged(auth, (currentUser) => {
			if (currentUser) {
				setUser(currentUser);
			} else {
				setUser(null);
			}
		});

		return () => unsubscribe();
	}, []);

	const handleGitHubLogin = async () => {
		try {
			// 永続化設定
			await setPersistence(auth, browserLocalPersistence);
			const result = await signInWithPopup(auth, provider);
			const user = result.user;
			console.log("Logged in user:", user);
			setUser(user);
		} catch (err) {
			console.error("GitHub Login Error:", err);
		}
	};

	useEffect(() => {
		if (user) {
			router.push("/home");
		}
	}, [user, router]);

	return (
		<div>
			<h1>Login with GitHub</h1>
			<button type="button" onClick={handleGitHubLogin}>
				Login with GitHub
			</button>
			{user && (
				<div>
					<p>{user.displayName}</p>
					<p>{user.email}</p>
				</div>
			)}
		</div>
	);
}
