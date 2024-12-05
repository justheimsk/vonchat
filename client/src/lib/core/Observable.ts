type SubscriptionCallback<T> = (arg: T) => void;

export class Subscription<T> {
  private observer: Observable<T>;
  public cb: SubscriptionCallback<T>;
  public id: string;
  
  public constructor(observer: Observable<T>, cb: SubscriptionCallback<T>) {
    this.observer = observer;
    this.cb = cb;
    this.id = crypto.randomUUID();
  }

  public unsubscribe() {
    return this.observer.unsubscribe(this.id);
  }
}

export class Observable<T> {
  private observers: Map<string, Subscription<T>>;

  public constructor() {
    this.observers = new Map();
  }

  public subscribe(cb: SubscriptionCallback<T>) {
   const subs = new Subscription(this, cb);
   this.observers.set(subs.id, subs);

   return subs;
  }

  public notify(arg: T) {
    for (const subs of this.observers.values()) {
      subs.cb(arg);
    }
  }

  public unsubscribe(id: string) {
    return this.observers.delete(id);
  }
}
