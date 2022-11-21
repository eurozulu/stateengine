# State Engine
An experimental play with a simple state engine.  

## Concept  
Based on a common example of a vending machine, a simple state engine has three states:  

- Empty
- Locked
- Unlocked

To tranistion through a state, an Action is applied.  
There are three Actions:  

- Push Button
- Insert Coin
- Add Bottle

The machine begins in the `Empty` state.  
In `Empty` state only the `Add Bottle` action will change the state to `Locked`.  
`Push button` and `Insert coin` have no effect, returning to Empty state.

In `Locked` state only the `Insert coin` action will move to `Unlocked`.  
`Push button` has no effect.  `Add Bottle` will increase the bottles but leave in
the `Locked` state.

In `Unlocked` state only `Push Button` action will move to `Locked` or remain in `Unlocked` depending on Credit.
If there are unused coins, the machine is held in `Unlocked`, otherwise it returns to `Locked`

`Insert coin` is rejected when the machine has run out of bottle.

### Design
The design is centered around the State.
State is a simple int/consts.

Each State has an actionMap of its own, defining which
actions it accepts and which state those actions will move to.
On an action received, the resulting mapped state becomes the next state.

Vending machine is built on top of the State.  This contains
the state of the Vending machine:
Its coins, bottle and current state.  
On each transition of state, the vending machine checks its own status
and may modify the transition of state accordingly.  

e.g. When `Push button` action is given in `Unlocked` State,
The usual transition is to return to `Locked`.
If vending machine still has unused coins, it prevents this,
holding the state of `Unlocked` until all the coins are exhausted.

