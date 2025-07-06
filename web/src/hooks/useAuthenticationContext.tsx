import { createContext, useContext, useEffect, useState } from "react";
import { useQuery } from "@tanstack/react-query";
import axios from "axios";
interface AuthenticationContextInterface {
    authenticated: boolean;
    setAuthenticated: (value: boolean) => void;
    isLoading: boolean;
}

export const AuthenticationContext = createContext<AuthenticationContextInterface | undefined>(undefined);

export function useAuthenticationContext() {
    const context = useContext(AuthenticationContext);
    if (context === undefined) {
        throw new Error("useCookieContext must be used within a CookieProvider");
    }
    return context
}

export function AtuhenticationProvider({ children }: { children: React.ReactNode }) {
    const [authenticated, setAuthenticated] = useState<boolean>(false);
    const { isLoading, isSuccess: sessionValid } = useQuery({
        queryKey: ["auth"],
        queryFn: async () => (await axios.get("/user/auth", {})).data,
        retry: false
    })

    useEffect(() => {
        setAuthenticated(sessionValid);
    }, [sessionValid]);

    return (
        <AuthenticationContext.Provider value={{ authenticated, setAuthenticated, isLoading }}>
            {children}
        </AuthenticationContext.Provider>
    );
}