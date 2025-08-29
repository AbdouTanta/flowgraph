// Function to get a specific cookie by name
export const getCookie = (name: string) => {
  const value = `; ${document.cookie}`;
  const parts = value.split(`; ${name}=`);
  if (parts.length === 2) {
    return parts.pop()?.split(";").shift();
  }
  return null;
};

// Function to set a specific cookie by name
export const setCookie = (
  name: string,
  value: string,
  days = 7,
  path = "/",
) => {
  const expires = new Date(Date.now() + days * 864e5).toUTCString();
  document.cookie = `${name}=${value}; expires=${expires}; path=${path}`;
};

// Function to decode JWT token (without verification)
// Note: This only decodes the payload, doesn't verify the signature
export const decodeJWT = (token: string) => {
  try {
    // JWT has 3 parts separated by dots: header.payload.signature
    const parts = token.split(".");
    if (parts.length !== 3) {
      throw new Error("Invalid JWT format");
    }

    // Decode the payload (second part)
    const payload = parts[1];

    // Add padding if needed for base64 decoding
    const paddedPayload = payload + "=".repeat((4 - (payload.length % 4)) % 4);

    // Decode base64
    const decodedPayload = atob(paddedPayload);

    // Parse JSON
    return JSON.parse(decodedPayload);
  } catch (error) {
    console.error("Error decoding JWT:", error);
    return null;
  }
};

// Function to get and decode JWT from cookies
export const getDecodedJWTFromCookie = (cookieName = "token") => {
  const token = getCookie(cookieName);
  if (!token) {
    return null;
  }

  return decodeJWT(token);
};

// Function to check if token is expired
export const isTokenExpired = (decodedToken: any) => {
  if (!decodedToken || !decodedToken.exp) {
    return true;
  }

  const currentTime = Math.floor(Date.now() / 1000);
  return decodedToken.exp < currentTime;
};
