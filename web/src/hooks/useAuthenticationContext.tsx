import { createContext, useContext, useEffect, useState } from "react";
import Cookies from "js-cookie";
import { useQuery } from "@tanstack/react-query";
import axios from "axios";

interface AuthenticationContextInterface {
    authenticated: boolean;
    setAuthenticated: (value: boolean) => void;
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
    const [sessionCookie, setSessionCookie] = useState<string | undefined>(undefined);

    useEffect(() => {
        const sessionCookie = Cookies.get("session");
        if (sessionCookie) {
            setSessionCookie(sessionCookie);
        }
    }, []);

    return (
        <AuthenticationContext.Provider value={{ authenticated, setAuthenticated }}>
            {children}
        </AuthenticationContext.Provider>
    );
}