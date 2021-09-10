1. 我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。
> 不应该 Wrap 这个 error给上层，原因。。。遇到sql.ErrNoRow这种情况，属于用户业务行为