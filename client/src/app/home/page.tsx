"use client";

import { auth } from "~/firebase/initializeApp";
import { useAuth } from "~/hooks/useAuth";
import Link from "next/link";
import { useEffect, useState } from "react";

export default function HomePage() {
	const { user } = useAuth();
	const [loading, setLoading] = useState<boolean>(true);
	const [error, setError] = useState<string | null>(null);
	// const { user } = useAuth();
	const [data, setData] = useState(null);

	console.log("user", user);

	useEffect(() => {
		if (!user) return;

		const fetchPythonData = async () => {
			setLoading(true);
			setError(null);

			const query = `
				query {
					me {
						id
						name
					}
				}
			`;

			try {
				const response = await fetch("http://localhost:8080/query", {
					method: "POST",
					headers: {
						"Content-Type": "application/json",
						Authorization: `Bearer ${await user?.getIdToken()}`, // トークンを使用
					},
					body: JSON.stringify({ query }),
				});

				if (!response.ok) {
					throw new Error(`HTTP error! status: ${response.status}`);
				}

				const result = await response.json();

				if (result.errors) {
					throw new Error(result.errors[0].message);
				}

				setData(result.data);
			} catch (err) {
				console.error(err);
			} finally {
				setLoading(false);
			}
		};

		fetchPythonData();
	}, [user]);

	const handleLogout = async () => {
		try {
			await auth.signOut();
			console.log("Logged out successfully!");
		} catch (error) {
			console.error("Error during logout:", error);
		}
	};

	return (
		<div className="min-h-screen bg-gradient-to-br from-blue-600 via-blue-400 to-yellow-400 text-white p-8">
			<header className="flex justify-between items-center bg-gray-900 bg-opacity-80 p-4 rounded-lg">
				<h1 className="text-3xl font-mono font-bold tracking-wide">
					Python World
				</h1>
				<p>{user?.email}</p>
				<button
					onClick={handleLogout}
					type="button"
					className="bg-green-300 text-gray-900 py-2 px-4 rounded-lg font-bold hover:bg-yellow-300 transition-transform transform hover:scale-105"
				>
					Logout
				</button>
			</header>
			<main className="mt-8">
				<div className="grid grid-cols-1 md:grid-cols-3 gap-8">
					<Link
						href="/"
						type="button"
						className="block text-center bg-green-300 hover:bg-blue-700 text-white py-6 px-8 rounded-xl shadow-lg text-lg font-bold tracking-wider uppercase transition-transform transform hover:scale-105"
					>
						Write Code
					</Link>
					<Link
						href="/"
						type="button"
						className="block text-center bg-green-300 hover:bg-gray-700 text-yellow-400 py-6 px-8 rounded-xl shadow-lg text-lg font-bold tracking-wider uppercase transition-transform transform hover:scale-105"
					>
						Run Script
					</Link>
					<Link
						href={`/${user?.uid}/create`}
						type="button"
						className="block text-center bg-green-300 hover:bg-yellow-500 text-gray-900 py-6 px-8 rounded-xl shadow-lg text-lg font-bold tracking-wider uppercase transition-transform transform hover:scale-105"
					>
						ペイモン作成
					</Link>
				</div>
			</main>
		</div>
	);
}
