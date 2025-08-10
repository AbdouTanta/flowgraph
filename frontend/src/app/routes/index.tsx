import { Redirect } from "wouter";

export default function Index() {
  // TODO - Redirect only if user is authenticated
  // If not authenticated, redirect to login page
  // If authenticated, redirect to flows page

  return <Redirect to="/flows" />;
}
