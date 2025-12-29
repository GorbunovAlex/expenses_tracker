// =============================================================================
// Auth Events - Global authentication event handling
// =============================================================================

type AuthEventListener = () => void;

class AuthEventEmitter {
  private listeners: Set<AuthEventListener> = new Set();

  /**
   * Subscribe to unauthorized (401) events
   */
  onUnauthorized(listener: AuthEventListener): () => void {
    this.listeners.add(listener);

    // Return unsubscribe function
    return () => {
      this.listeners.delete(listener);
    };
  }

  /**
   * Emit unauthorized event to all listeners
   */
  emitUnauthorized(): void {
    this.listeners.forEach((listener) => {
      try {
        listener();
      } catch (error) {
        console.error("Error in auth event listener:", error);
      }
    });
  }
}

// Singleton instance
export const authEvents = new AuthEventEmitter();
