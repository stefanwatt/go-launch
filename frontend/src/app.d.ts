export declare namespace GoLaunch {
  type GoMethod = "FuzzyFindDesktopEntry" | "LaunchApp";
  interface DesktopEntry {
    Id: string;
    Name: string;
    Exec: string;
    Icon: string;
  }

  interface GoMethodResponse {
    result: any;
    messageId: string;
  }

  interface GoMethodRequest {
    method: GoMethod;
    args: string; // adapt if they get more complex
    messageId: string;
  }
}
