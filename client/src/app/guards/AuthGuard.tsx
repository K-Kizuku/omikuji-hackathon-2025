"use client";

import { useRouter } from "next/navigation";
import { type ReactNode, useEffect } from "react";
import { Loading } from "~/components/Loading";
import { useAuth } from "~/hooks/useAuth";

export default function AuthGuard({ children }: { children: ReactNode }) {
	const { user, loading } = useAuth();
	const router = useRouter();

	useEffect(() => {
		if (!loading && !user) {
			router.push("/auth/login"); // 未認証ならログインページへ
		}
	}, [user, loading, router]);

	if (loading) {
		return <Loading />;
	}

	return <>{children}</>;
}
