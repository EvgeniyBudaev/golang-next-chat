export class BufferedWebSocket extends WebSocket {
  buffer: Event[] = [];
  
  private onMessageListeners: EventListenerOrEventListenerObject[] = [];

  constructor(url: string | URL, protocols?: string | string[] | undefined) {
    super(url, protocols);
    
    this.addEventListener<"message">("__message", this.onMessage.bind(this));
  }

  public addEventListener<K extends keyof WebSocketEventMap>(type: K | "__message", listener: (this: WebSocket, ev: WebSocketEventMap[K]) => any, options?: boolean | AddEventListenerOptions | undefined): void;
  public addEventListener(type: string, listener: EventListenerOrEventListenerObject, options?: boolean | AddEventListenerOptions | undefined): void;
  public addEventListener(type: unknown, listener: unknown, options?: unknown): void {
    if(type === "message") {
      this.onMessageListeners.push(listener as EventListenerOrEventListenerObject);
      this.flushBuffer();
      return;
    }

    if(type === "__message") {
      type = "message";
    }

    super.addEventListener(type as string, listener as EventListenerOrEventListenerObject, options as boolean | AddEventListenerOptions | undefined);
  }

  public removeEventListener<K extends keyof WebSocketEventMap>(type: K | "__message", listener: (this: WebSocket, ev: WebSocketEventMap[K]) => any, options?: boolean | EventListenerOptions | undefined): void;
  public removeEventListener(type: string, listener: EventListenerOrEventListenerObject, options?: boolean | EventListenerOptions | undefined): void;
  public removeEventListener(type: unknown, listener: unknown, options?: unknown): void {
    if(type === "message") {
      this.onMessageListeners = this.onMessageListeners.filter(l => l !== listener);
      return;
    }

    if(type === "__message") {
      type = "message";
    }

    super.removeEventListener(type as string, listener as EventListenerOrEventListenerObject, options as boolean | EventListenerOptions | undefined);
  }

  private flushBuffer() {
    if(this.onMessageListeners.length > 0) {
      this.buffer.forEach(message => {
        this.onMessageListeners.forEach(listener => "handleEvent" in listener ? listener.handleEvent(message) : listener(message));
      });

      this.buffer = [];
    }
  }

  private onMessage(message: MessageEvent) {
    if(this.onMessageListeners.length > 0) {
      this.onMessageListeners.forEach(listener => "handleEvent" in listener ? listener.handleEvent(message) : listener(message));
    } else {
      this.buffer.push(message);
    }
  }
}
