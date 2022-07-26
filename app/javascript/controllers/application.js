import { Application } from "@hotwired/stimulus"
import Alert from "./hello"

const application = Application.start()
Alert()

// Configure Stimulus development experience
application.debug = false
window.Stimulus   = application

export { application }
