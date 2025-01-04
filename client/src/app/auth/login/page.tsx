"use client";

import { useEffect, useState } from "react";
import {
	browserLocalPersistence,
	GithubAuthProvider,
	onAuthStateChanged,
	setPersistence,
	signInWithPopup,
	type User,
} from "firebase/auth";
import { auth, provider } from "~/firebase/initializeApp";
import { useRouter } from "next/navigation";
import { FaGithub } from "react-icons/fa6";

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

			// GitHubアクセストークンを取得
			const credential = GithubAuthProvider.credentialFromResult(result);
			const githubToken = credential?.accessToken;

			if (githubToken) {
				// アクセストークンをlocalStorageに保存
				localStorage.setItem("githubToken", githubToken);
				console.log("GitHub Token stored in localStorage:", githubToken);
			}

			setUser(user);
		} catch (err) {
			console.error("GitHub Login Error:", err);
		}
	};

	useEffect(() => {
		if (user) {
			router.push(`/${user.uid}`);
		}
	}, [user, router]);

	return (
		<div className="min-h-screen bg-gradient-to-br from-green-300 via-emerald-400 to-teal-500 flex items-center justify-center">
			<div className="w-96 bg-white shadow-lg rounded-xl p-6 text-center">
				<h1 className="text-2xl font-bold text-emerald-800 mb-4">
					🐍 Welcome to Pymon 🐍
				</h1>
				<p className="text-sm text-gray-600 mb-6">
					Sign in with GitHub to explore the world of Pymons!
				</p>
				<button
					type="button"
					onClick={handleGitHubLogin}
					className="w-full bg-emerald-600 relative text-white py-3 rounded-lg font-semibold hover:bg-emerald-700 transition-transform transform hover:scale-105"
				>
					<FaGithub size={24} className="absolute left-4" />
					Login with GitHub
				</button>
			</div>
		</div>
	);
}
