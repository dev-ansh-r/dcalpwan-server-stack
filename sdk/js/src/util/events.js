

class EventHandler {
  constructor() {
    this.EVENTS = Object.freeze({
      WARNING: 'warning',
      // Add more here as we go.
    })

    this.eventHandlers = {}
    this.dispatchEvent = (event, payload) => {
      if (this.eventHandlers[event]) {
        for (const handler of this.eventHandlers[event]) {
          handler(payload)
        }
      }
    }

    this.subscribe = (event, handler) => {
      if (!Object.values(this.EVENTS).includes(event)) {
        throw new Error(`Cannot subscribe to unsupported event type "${event}"`)
      }
      this.eventHandlers = {
        ...this.eventHandlers,
        [event]: this.eventHandlers[event] ? [...this.eventHandlers[event], handler] : [handler],
      }
    }

    this.unsubscribe = event => {
      if (!Object.values(this.EVENTS).includes(event)) {
        throw new Error(`Cannot unsubscribe from unsupported event type "${event}"`)
      }
      if (this.eventHandlers[event]) {
        delete this.eventHandlers[event]
      }
    }
  }
}

export default new EventHandler()
